import { useRouter } from 'next/router';
import { requests } from '../utils/requests';

interface SignInUserRequestBody {
  email?: string;
  password?: string;
}

interface UseSignInUserResponse {
  handleSignIn: (values: SignInUserRequestBody) => void;
}

const useSignInUser = (): UseSignInUserResponse => {
  const router = useRouter();
  const handleSignIn = async ({ email, password }: SignInUserRequestBody) => {
    try {
      await requests.post('/api/auth/login', { email, password });
      router.replace('/dashboard');
    } catch (error) {
      console.log('there was an error');
    }
  };

  return {
    handleSignIn,
  };
};

export default useSignInUser;
