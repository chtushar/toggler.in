import { useRouter } from 'next/router';
import { useMutation } from 'react-query';
import { queryKeys } from '../constants/queryKeys';
import { queryClient } from '../utils/requestClient';
import { requests } from '../utils/requests';

interface SignInUserRequestBody {
  email?: string;
  password?: string;
}

interface SignInUserResponseBody {
  user: {
    id?: string;
    firstName?: string;
    lastName?: string;
    email?: string;
  };
}

const useAuth = () => {
  const router = useRouter();
  const signIn = useMutation(
    ({
      email,
      password,
    }: SignInUserRequestBody): Promise<SignInUserResponseBody> =>
      requests.post('/api/auth/login', { email, password }),
    {
      onSuccess: (response) => {
        queryClient.setQueryData<SignInUserResponseBody['user']>(
          queryKeys.user,
          () => {
            return { ...response.user };
          },
        );
        router.replace('/dashboard');
      },
    },
  );

  const signOut = useMutation(() => requests.get('/api/auth/logout'), {
    onSuccess: () => {
      queryClient.clear();
      router.replace('/');
    },
  });

  return { signIn, signOut };
};

export default useAuth;
