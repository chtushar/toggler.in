import type { NextPage } from 'next';
import Layout from '../components/Layout';
import SignupForm from '../components/SignupForm';

const Signup: NextPage = () => {
  return (
    <Layout>
      <div>
        <div>
          <div>
            <SignupForm />
          </div>
        </div>
      </div>
    </Layout>
  );
};

export default Signup;
