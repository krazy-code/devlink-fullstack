import type { ProfileItem } from '@/services/profile/profile.types';
import { useMutation, useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router';
import type { AuthFormBody } from '../../services/auth/auth.types';
import networkGenerator from '../../services/network-generator';

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
      navigate('/', { replace: true });
    },
  });

  const logout = useMutation({
    mutationFn: () => networkGenerator(`${apiPrefix}/logout`, 'post'),
    onSuccess: () => {
      window.localStorage.clear();
      navigate('/', { replace: true });
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
