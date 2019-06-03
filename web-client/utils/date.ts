type DateType = Date | string;

function assertDate(date: DateType): Date {
  return date instanceof Date ? date : new Date(date);
}

const WEEKDAYS = [
  'Sunday',
  'Monday',
  'Tuesday',
  'Wednesday',
  'Thursday',
  'Friday',
  'Saturday',
];

export function formatDate(date: DateType): string {
  date = assertDate(date);
  const d = date.getDate();
  const m = date.getMonth() + 1;
  const y = date.getFullYear();

  const weekday = WEEKDAYS[date.getDay()];

  return `${weekday}, ${d}.${m}.${y}`;
}

export function formatDateTime(date: DateType): string {
  date = assertDate(date);
  const h = date.getHours();
  const m = date.getMinutes();

  return `${formatDate(date)} ${h}:${m}`;
}

export function sameDay(d1: Date, d2: Date): boolean {
  return (
    d1.getFullYear() === d2.getFullYear() &&
    d1.getMonth() === d2.getMonth() &&
    d1.getDate() === d2.getDate()
  );
}
