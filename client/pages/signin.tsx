import { Box, Grid, Flex, Button } from '../ds';
import type { NextPage } from 'next';
import Layout from '../components/Layout';
import SignInForm from '../components/SignInForm';

const SignIn: NextPage = () => {
  return (
    <Layout>
      <Box css={{ width: '100%' }}>
        <Grid flow='row' columns='5'>
          <Box
            css={{
              gridColumn: '1 / 3',
              '@bp2': {
                gridColumn: '1 / 4',
              },
              '@bp3': {
                gridColumn: '1 / 6',
              },
            }}
          >
            <SignInForm />
          </Box>
        </Grid>
      </Box>
    </Layout>
  );
};

export default SignIn;
