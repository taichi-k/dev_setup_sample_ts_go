import type { Meta, StoryObj } from '@storybook/react';
import Button from './Button';

const meta: Meta<typeof Button> = {
  component: Button,
  title: 'UI/Button',
  args: {
    label: 'Button',
  },
};
export default meta;

type Story = StoryObj<typeof Button>;

export const Primary: Story = {
  args: {
    label: 'Primary',
  },
};
export const Secondary: Story = {
  args: {
    label: 'Secondary',
  },
};
export const WithIcon: Story = {
  args: {
    label: 'With Icon',
  },
};
