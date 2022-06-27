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
  data: {
    id?: string;
    name?: string;
    email?: string;
    emailVerified?: boolean;
    created_at?: string;
    updated_at?: string;
  };
}

const useAuth = () => {
  const router = useRouter();
  const signIn = useMutation(
    ({
      email,
      password,
    }: SignInUserRequestBody): Promise<SignInUserResponseBody> =>
      requests.post('v1/auth/signin', { email, password }),
    {
      onSuccess: (response) => {
        queryClient.setQueryData<SignInUserResponseBody['data']>(
          queryKeys.user,
          () => {
            return { ...response.data };
          },
        );
        router.replace('/dashboard');
      },
    },
  );

  const signOut = useMutation(() => requests.get('v1/auth/signout'), {
    onSuccess: () => {
      queryClient.clear();
      router.replace('/');
    },
  });

  return { signIn, signOut };
};

export default useAuth;
