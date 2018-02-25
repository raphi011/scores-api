import { formatDate } from './dateFormat';

const date = new Date(2017, 0, 1);

test('formatDate', () => {
  expect(formatDate(date)).toBe('1.1.2017');
});
