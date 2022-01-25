import Layout from '../../components/Layout';
import Sidebar from '../../components/Sidebar';
import FeatureFlags from '../../components/FeatureFlags';
import { Grid } from '../../ds';
import { requests } from '../../utils/requests';

const Dashboard = () => {
  return (
    <Layout>
      <Grid css={{ height: '100%' }} columns={4} gap={4}>
        <Sidebar />
        <FeatureFlags />
      </Grid>
    </Layout>
  );
};

export const getServerSideProps = async (ctx: any) => {
  try {
    const data = await requests.post('/api/flag', undefined, {
      headers: ctx?.req?.headers?.cookie
        ? { cookie: ctx.req.headers.cookie }
        : undefined,
    });

    return {
      props: {
        data,
      },
    };
  } catch (error) {
    return {
      redirect: {
        destination: '/signin',
        permanent: false,
      },
    };
  }
};

export default Dashboard;
