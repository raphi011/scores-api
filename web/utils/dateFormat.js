export function formatDate(date) {
  const d = date.getDate();
  const m = date.getMonth() + 1;
  const y = date.getFullYear();

  return `${d}.${m}.${y}`;
}

export function formatDateTime(date) {
  const h = date.getHours();
  const m = date.getMinutes();

  return `${formatDate(date)} ${h}:${m}`;
}
