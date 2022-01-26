import { Button } from './Button';
import { Icon } from './Icon';
import { useSelect } from 'downshift';
import { styled } from '../stitches.config';
import { Flex } from './Flex';

interface SelectProps {
  items: Array<Record<string, any>>;
}

const SelectWrapper = styled('div', {
  position: 'relative',
});

const List = styled(Flex, {
  position: 'absolute',
  listStyle: 'none',
  marginTop: '$2',
  borderRadius: '$space$2',
  overflow: 'hidden',
  filter: 'drop-shadow(0px 4px 10px rgba(0, 0, 0, 0.1))',
  zIndex: 10,
  width: '100%',
});

const ListItem = styled('li', {
  background: '$slate1',
  padding: '$4',
  fontSize: '$14',
  variants: {
    selected: {
      true: {
        background: '$slate3',
      },
      false: {
        background: '$slate1',
      },
    },
  },
});

export const Select = ({ items }: SelectProps) => {
  const {
    isOpen,
    highlightedIndex,
    getToggleButtonProps,
    getMenuProps,
    getItemProps,
    selectedItem,
  } = useSelect({
    items,
  });
  return (
    <SelectWrapper>
      <Button type='button' variant='select' {...getToggleButtonProps()}>
        {selectedItem || 'Choose an option'}{' '}
        <Icon
          className={isOpen ? 'ri-arrow-up-s-line' : 'ri-arrow-down-s-line'}
        />
      </Button>
      <List as='ul' direction='column' align='stretch' {...getMenuProps()}>
        {isOpen &&
          items.map(({ value, label }, index) => (
            <ListItem
              key={value}
              selected={highlightedIndex === index}
              {...getItemProps({ item: value, index })}
            >
              {label}
            </ListItem>
          ))}
      </List>
    </SelectWrapper>
  );
};

export default Select;
