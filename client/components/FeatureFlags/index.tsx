import { Flex, Card } from '../../ds';

const FeatureFlags = () => {
  return (
    <Flex
      direction='column'
      align='stretch'
      gap={10}
      css={{ gridColumn: '2 / 5' }}
    >
      <Card variant='primary2'>Dashboard</Card>
      <Card variant='primary1'>Dashboard</Card>
    </Flex>
  );
};

export default FeatureFlags;
