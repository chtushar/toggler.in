import { requests } from '../utils/requests';

interface SignInUserRequestBody {
  email?: string;
  password?: string;
}

interface UseSignInUserResponse {
  handleSignIn: (values: SignInUserRequestBody) => void;
}

const useSignInUser = (): UseSignInUserResponse => {
  const handleSignIn = async ({ email, password }: SignInUserRequestBody) => {
    await requests.post('/api/auth/login', { email, password });
  };

  return {
    handleSignIn,
  };
};

export default useSignInUser;
