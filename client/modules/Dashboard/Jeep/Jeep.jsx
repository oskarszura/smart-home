// @flow
import _ from 'lodash';
import React from 'react';
import { withRouter } from 'react-router';
import * as agentsTypes from 'client/models/agents/types';
import Joystick from 'client/components/Joystick';
import Switch from 'client/components/Switch';
import * as proxyTypes from 'client/models/proxy/types';
import * as proxyConstants from 'client/models/proxy/constants';

type Props = {
  agent: agentsTypes.Agent,
  onPositionChange: (agentsTypes.Agent, { left: number, top: number }) => void,
  onToggle: () => void,
  setup: () => void,
  status: proxyTypes.Status,
};

class Jeep extends React.PureComponent<Props> {
  render() {
    const { agent, onPositionChange, onToggle, status } = this.props;

    const isConnected = status === proxyConstants.STATUS_CONNECTED;
    const isPending = status === proxyConstants.STATUS_PENDING;

    return (
      <div className="jeep-panel">
        <div className="jeep-panel__section">
          <div className="c-control">
            Device connection
            <div className="c-control__content">
              <Switch className="" isOn={isConnected} onToggle={_.partial(onToggle, agent, isConnected)} />
            </div>
          </div>
        </div>
        <div className="jeep-panel__section">
          { !isPending &&
            <Joystick
              isEnabled={isConnected}
              onPositionChange={(left: number, top: number) => {
                onPositionChange(agent, { left, top, flag: null });
              }}
            />
          }
          { isPending &&
            <div className="c-loader" />
          }
        </div>
      </div>
    );
  }
}

export default withRouter(Jeep);
