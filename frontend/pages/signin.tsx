import type { NextPage } from 'next';
import Layout from '../components/Layout';
import SignInForm from '../components/SignInForm';

const SignIn: NextPage = () => {
  return (
    <Layout>
      <div>
        <div>
          <div>
            <SignInForm />
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default SignIn;
