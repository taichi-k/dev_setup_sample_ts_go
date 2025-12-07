import { expect, test } from '@playwright/test';

test('test', async ({ page }) => {
  await page.goto('http://localhost:5173/');
  await page.getByRole('button', { name: 'HEADER' }).click();
  await page.getByRole('heading', { name: 'Go + React Monorepo' }).click();
  await page.getByText('User 1').click();
  await page.getByRole('heading', { name: 'Go + React Monorepo' }).click();
  await expect(page.getByRole('heading')).toMatchAriaSnapshot(`- heading "Go + React Monorepo" [level=1]`);
});
