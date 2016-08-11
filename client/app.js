import riot from 'riot';
import reduxMixin from 'riot-redux-mixin';

import './components/instancelist.tag';
import './components/memorycontrol.tag';
import './components/diskcontrol.tag';
import './components/networkcontrol.tag';
import './components/instancecontrols.tag';

import storeFactory from './store';
import * as wares from './store/middleware';
import { setInstances } from './store/action-creators';

import TaskMonitor from './task-monitor';
import TrafficSimulator from './traffic-simulator';

var taskMonitor = new TaskMonitor();
var trafficSim = new TrafficSimulator()
var store = storeFactory(wares.logging, wares.traffic(trafficSim), wares.remoteRelay(() => taskMonitor.update()));

trafficSim.init(store);
taskMonitor.init(store);
taskMonitor.start();

riot.mixin('redux', reduxMixin(store));
riot.mount('*', {});
