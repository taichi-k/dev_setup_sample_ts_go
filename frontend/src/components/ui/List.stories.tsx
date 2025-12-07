import type { Meta, StoryObj } from '@storybook/react';
import { mockUsers } from '../../__mock__/user';
import List from './list';

const meta: Meta<typeof List> = {
  component: List,
  title: 'Components/UI/List',
  tags: ['autodocs'],
  args: {
    items: mockUsers.map((user) => user.name),
  },
};

export default meta;
type Story = StoryObj<typeof List>;

export const Default: Story = {};
