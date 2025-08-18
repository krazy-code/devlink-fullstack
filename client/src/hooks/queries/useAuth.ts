import type { ProfileItem } from '@/services/profile/profile.types';
import { useMutation, useQuery } from '@tanstack/react-query';
import { useNavigate } from '@tanstack/react-router';
import networkGenerator from '../../services/_config/network-generator';
import type { AuthFormBody } from '../../services/auth/auth.types';

const apiPrefix = '/auth';

export default function useAuth() {
  const navigate = useNavigate();

  const register = useMutation({
    mutationFn: (body: AuthFormBody) =>
      networkGenerator<unknown, AuthFormBody>(
        `${apiPrefix}/register`,
        'post',
        body
      ),
  });

  const login = useMutation({
    mutationFn: (body: AuthFormBody) =>
      networkGenerator<{ token: string }, AuthFormBody>(
        apiPrefix,
        'post',
        body
      ),
    onSuccess: (res) => {
      const token = res.token;
      window.localStorage.setItem('token', token);
      navigate({ to: '/feed', replace: true });
    },
  });

  const logout = useMutation({
    mutationFn: () => networkGenerator(`${apiPrefix}/logout`, 'post'),
    onSuccess: () => {
      window.localStorage.clear();
      navigate({ to: '/', replace: true });
    },
  });

  const useGetProfile = () =>
    useQuery({
      queryKey: ['profile'],
      queryFn: () =>
        networkGenerator<ProfileItem>(`${apiPrefix}/profile`, 'get'),
    });

  return {
    register,
    login,
    logout,
    getProfile: useGetProfile,
  };
}
