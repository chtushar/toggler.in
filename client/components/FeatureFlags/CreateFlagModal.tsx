import { Dialog, Button, Icon } from '../../ds';
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
        <Dialog.Content>Yes</Dialog.Content>
      </Dialog.Portal>
    </Dialog.Root>
  );
};
