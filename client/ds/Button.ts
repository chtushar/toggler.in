import { styled } from '../stitches.config';

export const Button = styled('button', {
  cursor: 'pointer',
  padding: '$4 $8',
  backgroundColor: '$slate12',
  color: '$slate1',
  border: 'none',
  borderRadius: '$space$4',
  '&:disabled': {
    cursor: 'default',
    backgroundColor: '$slate8',
  },
});
