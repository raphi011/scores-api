import { formatDate } from './date';

const date = new Date(2017, 0, 1);

test('formatDate', () => {
  expect(formatDate(date)).toBe('Sunday, 1.1.2017');
});
