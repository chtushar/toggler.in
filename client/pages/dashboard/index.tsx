import Layout from '../../components/Layout';
import { requests } from '../../utils/requests';

const Dashboard = () => {
  return <Layout>Dashboard</Layout>;
};

export const getServerSideProps = async (ctx: any) => {
  const data = await requests.post('/api/flag', undefined, {
    headers: ctx?.req?.headers?.cookie
      ? { cookie: ctx.req.headers.cookie }
      : undefined,
  });

  return {
    props: {
      fallback: {
        '/api/flag': data,
      },
    },
  };
};

export default Dashboard;
