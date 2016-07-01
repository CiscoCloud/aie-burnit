import riot from 'riot';
import reduxMixin from 'riot-redux-mixin';

import './components/instancelist.tag';
import './components/memorycontrol.tag';
import './components/instancecontrols.tag';

import store from './store';
import { setInstances } from './store/action-creators';
import StatusChecker from './status-check';

var statusCheck = new StatusChecker(store);
statusCheck.start();

riot.mixin('redux', reduxMixin(store));
riot.mount('*', {});
