import { Box, Grid, Flex, Button } from '../ds';
import type { NextPage } from 'next';
import Layout from '../components/Layout';
import SignupForm from '../components/SignupForm';

const Signup: NextPage = () => {
  return (
    <Layout>
      <Box css={{ width: '100%' }}>
        <Grid flow='row' columns='5'>
          <Box css={{ gridColumn: '1 / 3' }}>
            <SignupForm />
          </Box>
        </Grid>
      </Box>
    </Layout>
  );
};

export default Signup;
