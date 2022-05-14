import { CreateFlagModal } from './CreateFlagModal';
import { Flex, Card, Text, Button, Icon, Dialog } from '../../ds';

const FeatureFlags = () => {
  return (
    <Flex
      direction='column'
      align='stretch'
      gap={10}
      css={{ gridColumn: '2 / 5' }}
    >
      <Card
        direction='row'
        align='center'
        justify='between'
        variant='primary2'
        padding={32}
      >
        <Text size={48} weight='bold' color='slate12'>
          Feature Flags
        </Text>
        <CreateFlagModal />
      </Card>
      <Card variant='primary1'>Dashboard</Card>
    </Flex>
  );
};

export default FeatureFlags;
