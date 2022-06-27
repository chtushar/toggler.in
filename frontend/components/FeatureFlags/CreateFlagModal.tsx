import {
  Dialog,
  Button,
  Flex,
  Icon,
  Select,
  Text,
  TextField,
  TextArea,
  FieldSet,
} from '../../ds';
import React from 'react';

export const CreateFlagModal = () => {
  return (
    <Dialog.Root>
      <Button as={Dialog.Trigger} appearance='primary'>
        <Icon className='ri-add-line' />
        Add a new flag
      </Button>
      <Dialog.Portal>
        <Dialog.Overlay />
        <Flex as={Dialog.Content} direction='column' gap={8}>
          <Dialog.Title>Create a new flag</Dialog.Title>
          <Flex as='form' direction='column' gap={8} align='stretch'>
            <Flex direction='row' gap={6}>
              <FieldSet>
                <Text as='label' size={14} color='slate10' weight='semiBold'>
                  Team
                </Text>
                <Select
                  items={[{ value: 'devfolio', label: 'Devfolio' }]}
                  selectedOption={{ value: 'devfolio', label: 'Devfolio' }}
                />
              </FieldSet>
              <FieldSet>
                <Text as='label' size={14} color='slate10' weight='semiBold'>
                  Name
                </Text>
                <TextField placeholder='Flag Name' />
              </FieldSet>
              <FieldSet>
                <Text as='label' size={14} color='slate10' weight='semiBold'>
                  Key
                </Text>
                <TextField placeholder='Flag Key' />
              </FieldSet>
            </Flex>
            <Flex direction='row' gap={6}>
              <FieldSet>
                <Text as='label' size={14} color='slate10' weight='semiBold'>
                  Description
                </Text>
                <TextArea placeholder='Flag Name' />
              </FieldSet>
              <FieldSet>
                <Text as='label' size={14} color='slate10' weight='semiBold'>
                  Type
                </Text>
                <Select
                  items={[{ value: 'boolean', label: 'Boolean' }]}
                  selectedOption={{ value: 'boolean', label: 'Boolean' }}
                />
              </FieldSet>
            </Flex>
            <Flex direction='row' justify='end' gap={2}>
              <Button appearance='secondary' type='button' as={Dialog.Close}>
                Cancel
              </Button>
              <Button type='submit' appearance='primary'>
                Create
              </Button>
            </Flex>
          </Flex>
        </Flex>
      </Dialog.Portal>
    </Dialog.Root>
  );
};
