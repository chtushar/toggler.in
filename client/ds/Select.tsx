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
});

const ListItem = styled('li', {});

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
      <List as='ul' direction='column' {...getMenuProps()}>
        {isOpen &&
          items.map(({ value, label }, index) => (
            <ListItem key={value} {...getItemProps({ item: value, index })}>
              {label}
            </ListItem>
          ))}
      </List>
    </SelectWrapper>
  );
};

export default Select;
