import { Box, Grid, Flex, Button } from '../ds';
import type { NextPage } from 'next';
import Layout from '../components/Layout';
import SignInForm from '../components/SignInForm';
import { requests } from '../utils/requests';

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

export const getServerSideProps = async (ctx: any) => {
  try {
    const { isLoggedIn } = await requests.get('/api/user/status', {
      headers: ctx?.req?.headers?.cookie
        ? { cookie: ctx.req.headers.cookie }
        : undefined,
    });

    if (isLoggedIn) {
      return {
        redirect: {
          destination: '/dashboard',
          permanent: false,
        },
      };
    }
    return { props: {} };
  } catch (error) {
    return {
      props: {},
    };
  }
};

export default SignIn;
