import { requests } from '../utils/requests';

interface CreateUserRequestBody {
  firstName?: string;
  lastName?: string;
  email?: string;
  password?: string;
}

interface UseCreateUserResponse {
  handleCreateUser: (values: CreateUserRequestBody) => void;
}

const useCreateUser = (): UseCreateUserResponse => {
  const handleCreateUser = async ({
    firstName,
    lastName,
    email,
    password,
  }: CreateUserRequestBody) => {
    await requests.post('/api/user', { firstName, lastName, email, password });
  };

  return {
    handleCreateUser,
  };
};

export default useCreateUser;
