import { test } from '@playwright/test';

test('ホーム画面のスクリーンショット', async ({ page }) => {
  await page.goto('http://localhost:5173/'); // or your URL

  await page.screenshot({
    path: 'vrt_actual/home.png', // reg-suit の actualDir と合わせる
    fullPage: true,
  });
});
