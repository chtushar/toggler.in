import { styled } from '../stitches.config';

export const Button = styled('button', {
  cursor: 'pointer',
  padding: '$4 $8',
  backgroundColor: '$gray12',
  color: '$gray1',
  border: 'none',
  borderRadius: '$space$1',
  '&:disabled': {
    cursor: 'default',
    backgroundColor: '$gray8',
  },
});
