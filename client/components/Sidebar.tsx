import { Card, Flex, Text, Icon, Button, Select } from '../ds';

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
          <Select
            items={[{ value: 'devfolio', label: 'Devfolio' }]}
            selectedOption={{ value: 'devfolio', label: 'Devfolio' }}
          />
        </Flex>
        <Flex direction='column' gap={2}>
          <Text as='label' size={14} color='slate10' weight='semiBold'>
            Mode
          </Text>
          <Select
            items={[
              { value: 'development', label: 'Development' },
              { value: 'production', label: 'Production' },
            ]}
          />
        </Flex>
      </Card>
      <Card direction='column' align='stretch' gap={2} variant='primary2'>
        <Button variant='menu-button'>
          <Icon className='ri-toggle-line' />
          Feature Flags
        </Button>
        <Button variant='menu-button'>
          <Icon className='ri-team-line' />
          Team
        </Button>
      </Card>
      <Card direction='column' align='stretch' gap={2} variant='primary2'>
        <Button variant='transparent'>
          <Icon className='ri-settings-3-line' />
          Settings
        </Button>
      </Card>
    </Flex>
  );
};

export default Sidebar;
