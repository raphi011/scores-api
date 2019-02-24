import React from 'react';

interface Props<T> {
  groupItems: (items: T[]) => T[][];
  renderHeader: (item: T[]) => JSX.Element;
  renderList: (items: T[]) => JSX.Element;
  items: T[];
}

export default function GroupedList<T>({
  groupItems,
  items,
  renderHeader,
  renderList,
}: Props<T>) {
  const groupedItems = groupItems(items);

  const groupsWithHeaders: JSX.Element[] = [];

  groupedItems.forEach(group => {
    groupsWithHeaders.push(renderHeader(group));
    groupsWithHeaders.push(renderList(group));
  });

  return <>{groupsWithHeaders}</>;
}
