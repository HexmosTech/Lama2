import { memo, useEffect } from 'react';
import { Column } from '@firecamp/ui';
import cx from 'classnames';
import { shallow } from 'zustand/shallow';
import { ErrorBoundary } from 'react-error-boundary';
import Home from './home/Home';
import TabContainerBodyRequest from './TabContainerBodyRequest';
import ErrorPopup from '../common/error-boundary/ErrorPopup';
import { ITabStore, useTabStore } from '../../store/tab';
import { ERequestTypes } from '@firecamp/types';
import { ETabEntityTypes } from './types/tab';
const TabContainerBody = () => {
  const activeTab = useTabStore((s: ITabStore) => s.activeTab, shallow);
  const { list: tabs, orders } = useTabStore.getState();
  const { open: openTab } = useTabStore.getState();

  const _openTab = (type?: ERequestTypes | 'environment') => {
    openTab({ __meta: { type } }, { id: '', type: ETabEntityTypes.Request });
  };
  useEffect(() => {
    _openTab(ERequestTypes.Rest);
  }, []);
  return (
    <Column flex={1} className="invisible-scrollbar overflow-hidden">
      <div className={cx('fc-h-full invisible-scrollbar tab-content')}>
        <div
          className={cx('tab-pane', {
            active: activeTab == 'home',
          })}
        >
          {/* <Home /> */}
        </div>
        {orders.map((tabId, i) => (
          <div
            className={cx('tab-pane', {
              active: activeTab == tabId,
            })}
            key={tabId}
          >
            <ErrorBoundary
              FallbackComponent={ErrorPopup}
              // onError={(error, info) => {
              //   console.log({ error, info });
              //   close.byIds([tabObj.id]);
              //   changeActiveTab('home');
              // }}
            >
              <TabContainerBodyRequest tab={tabs[tabId]} key={tabId} />
            </ErrorBoundary>
          </div>
        ))}
      </div>
    </Column>
  );
};

export default memo(TabContainerBody);
