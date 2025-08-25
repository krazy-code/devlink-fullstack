/* eslint-disable @typescript-eslint/no-explicit-any */
import { useCallback, useMemo } from 'react';

type ToastStatus = 'success' | 'error' | 'info';
type ToastFn = (opts: {
  title?: string;
  description?: string;
  status?: ToastStatus;
}) => void;

type SetFieldErrorFn = (
  name: string,
  error: { type?: string; message?: string }
) => void;

interface ErrorPayload {
  status?: number;
  message?: string;
  errors?: Record<string, string | string[]>;
}

/**
 * useResponse - helper hook to show toast notifications and set form field errors from server responses.
 *
 * - Accepts an optional toast function (e.g. from your UI library). If not provided it will fallback to window.alert.
 * - handleSuccess shows a success toast.
 * - handleError parses common error shapes (Axios-like or plain) and:
 *     - if 400 + errors present and setError provided: calls setError for each field
 *     - otherwise shows an error toast with a message
 */
export default function useResponse(providedToast?: ToastFn) {
  const toast: ToastFn = useMemo(
    () =>
      providedToast ??
      ((opts) => {
        const text = [opts.title, opts.description].filter(Boolean).join(' - ');
        // Minimal fallback toast
        if (opts.status === 'error') {
          window.alert(`Error: ${text || 'Something went wrong'}`);
        } else {
          window.alert(text || 'Done');
        }
      }),
    [providedToast]
  );

  const showToast = useCallback(
    (opts: { title?: string; description?: string; status?: ToastStatus }) => {
      try {
        toast(opts);
      } catch {
        // swallow toast errors to avoid crashing UI flow
        // fallback to alert
        const text = [opts.title, opts.description].filter(Boolean).join(' - ');
        window.alert(text || (opts.status === 'error' ? 'Error' : 'Info'));
      }
    },
    [toast]
  );

  const setFieldErrorsFromPayload = useCallback(
    (errors: Record<string, any>, setError?: SetFieldErrorFn) => {
      if (!setError || !errors) return;
      Object.entries(errors).forEach(([field, value]) => {
        const message = Array.isArray(value)
          ? value.join(' ')
          : String(value ?? '');
        setError(field, { type: 'server', message });
      });
    },
    []
  );

  const handleSuccess = useCallback(
    (message?: string, description?: string) => {
      showToast({
        title: message ?? 'Success',
        description,
        status: 'success',
      });
    },
    [showToast]
  );

  const handleError = useCallback(
    (err: any, setError?: SetFieldErrorFn) => {
      // Support common shapes: Axios (err.response.data), fetch (err?.data), or plain error
      const payload: ErrorPayload =
        err?.response?.data ?? err?.data ?? err ?? {};
      const status = err?.response?.status ?? payload?.status;

      // Bad Request with field errors
      if (status === 400 && payload?.errors && setError) {
        setFieldErrorsFromPayload(payload.errors, setError);
        showToast({
          title: payload.message ?? 'Validation error',
          status: 'error',
        });
        return;
      }

      // Generic error message
      const message =
        payload?.message ?? err?.message ?? 'Something went wrong';
      showToast({ title: message, status: 'error' });
    },
    [setFieldErrorsFromPayload, showToast]
  );

  return {
    handleSuccess,
    handleError,
    setFieldErrorsFromPayload,
  };
}
