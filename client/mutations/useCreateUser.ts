import { requests } from '../utils/requests';

interface CreateUserRequestBody {
  email: string;
  password: string;
}

interface UseCreateUserResponse {
  handleCreateUser: (values: CreateUserRequestBody) => void;
}

const useCreateUser = (): UseCreateUserResponse => {
  const handleCreateUser = async ({
    email,
    password,
  }: CreateUserRequestBody) => {
    await requests.post('/api/user', { email, password });
  };

  return {
    handleCreateUser,
  };
};

export default useCreateUser;
