import React, { ReactNode } from 'react';

interface IProps<T> {
  groupItems: (items: T[]) => T[][];
  renderHeader: (T) => ReactNode;
  renderList: (items: T[]) => ReactNode;
  items: T[];
}

export default function GroupedList<T>({
  groupItems,
  items,
  renderHeader,
  renderList,
}: IProps<T>) {
  const groupedItems = groupItems(items);

  const groupsWithHeaders = [];

  groupedItems.forEach(group => {
    groupsWithHeaders.push(renderHeader(group));
    groupsWithHeaders.push(renderList(group));
  });

  return <>{groupsWithHeaders}</>;
}
