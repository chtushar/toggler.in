import { styled } from '../stitches.config';
import { Flex } from './Flex';

export const Card = styled(Flex, {
  appearance: 'none',
  border: 'none',
  boxSizing: 'border-box',
  font: 'inherit',
  lineHeight: '1',
  outline: 'none',
  padding: '$4',
  textAlign: 'inherit',
  verticalAlign: 'middle',
  WebkitTapHighlightColor: 'rgba(0, 0, 0, 0)',

  textDecoration: 'none',
  color: 'inherit',
  flexShrink: 0,
  borderRadius: '$space$4',
  position: 'relative',

  variants: {
    variant: {
      primary2: {
        backgroundColor: '$primary2',
      },
      primary1: {
        backgroundColor: '$primary1',
      },
    },
  },
});
