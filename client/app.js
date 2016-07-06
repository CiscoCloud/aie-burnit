import riot from 'riot';
import reduxMixin from 'riot-redux-mixin';

import './components/instancelist.tag';
import './components/memorycontrol.tag';
import './components/instancecontrols.tag';

import storeFactory from './store';
import { setInstances } from './store/action-creators';
import TaskMonitor from './task-monitor';
import * as wares from './store/middleware';

var taskMonitor = new TaskMonitor();
var store = storeFactory(wares.logging, wares.remoteRelay(() => taskMonitor.update()));

taskMonitor.init(store);
taskMonitor.start();

riot.mixin('redux', reduxMixin(store));
riot.mount('*', {});
