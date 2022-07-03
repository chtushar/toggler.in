import { useRouter } from 'next/router';
import { useMutation } from 'react-query';
import { queryKeys } from '../constants/queryKeys';
import { queryClient } from '../utils/requestClient';
import { requests } from '../utils/requests';

interface CreateUserRequestBody {
  name?: string;
  email?: string;
  password?: string;
}

interface CreateUserResponseBody {
  timestamp: string;
  uri: string;
  success: boolean;
  user: {
    id?: string;
    firstName?: string;
    lastName?: string;
    email?: string;
  } | null;
  error: {} | null;
}

const useCreateUser = () => {
  const router = useRouter();
  const mutation = useMutation(
    ({
      name,
      email,
      password,
    }: CreateUserRequestBody): Promise<CreateUserResponseBody> =>
      requests.post('v1/users/signup', {
        name,
        email,
        password,
      }),
    {
      onSuccess: (response) => {
        queryClient.setQueryData<CreateUserResponseBody['user']>(
          queryKeys.user,
          () => {
            return { ...response.user };
          },
        );
        router.replace('/dashboard');
      },
    },
  );

  return mutation;
};

export default useCreateUser;
