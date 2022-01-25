import { Card, Flex, Text } from '../ds';

const Sidebar = (): JSX.Element => {
  return (
    <Flex
      direction='column'
      align='stretch'
      gap={4}
      css={{
        gridColumn: '1 / 2',
      }}
    >
      <Card variant='primary2' direction='column' gap={4}>
        <Flex direction='column' gap={2}>
          <Text as='label' size={14} color='slate10' weight='semiBold'>
            Team
          </Text>
          <Text as='p' size={18} color='slate12' weight='semiBold'>
            Devfolio
          </Text>
        </Flex>
        <Flex direction='column' gap={2}>
          <Text as='label' size={14} color='slate10' weight='semiBold'>
            Team
          </Text>
          <Text as='p' size={18} color='slate12' weight='semiBold'>
            Devfolio
          </Text>
        </Flex>
      </Card>
      <Card variant='primary2'>Dashboard</Card>
      <Card variant='primary2'>Dashboard</Card>
    </Flex>
  );
};

export default Sidebar;
