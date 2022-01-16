import Layout from '../../components/Layout';
import { requests } from '../../utils/requests';

const Dashboard = () => {
  return <Layout>Dashboard</Layout>;
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
