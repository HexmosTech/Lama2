import { FC } from "react";
import { File, Github } from 'lucide-react';
import { SiDiscord } from '@react-icons/all-files/si/SiDiscord';

import { IHelp } from './interfaces/Help.interfaces'

const HELPDESK_TYPES = {
  DOCUMENTATION: 'documentation',
  GITHUB: 'github',
  JOIN_DISCORD: 'join_discord'
};

const helpDesks = {
  [HELPDESK_TYPES.DOCUMENTATION]: {
    id: HELPDESK_TYPES.DOCUMENTATION,
    name: 'Lama2 Docs',
    link: 'https://hexmos.com/lama2/index.html',
    image: <File size={24} strokeWidth={1.5} />
  },

};

/**
 * Firecamp help support with Doc, GitHub, and Discord.
 */
const Help: FC<IHelp> = ({ docLink = '', client = 'http' }) => {
  let _renderLink = (type = '') => {
    switch (type) {
      case HELPDESK_TYPES.DOCUMENTATION:
        return (
          <a
            className="text-app-foreground-inactive"
            href={docLink || 'https://hexmos.com/lama2/index.html'}
            target={'_blank'}
          >
            {helpDesks[type] ? helpDesks[type].name : ''}
          </a>
        );
        break;
      case HELPDESK_TYPES.GITHUB:
        return (
          <a
            className="text-app-foreground-inactive"
            href={
              client
                ? `https://github.com/firecamp-dev/firecamp/issues/new?assignees=&labels=&template=bug_report.md&title=[${client}]%20Title%20or%20Feature%20request`
                : helpDesks[type]
                  ? helpDesks[type].link
                  : ''
            }
            target={'_blank'}
          >
            {helpDesks[type] ? helpDesks[type].name : ''}
          </a>
        );
        break;
      case HELPDESK_TYPES.JOIN_DISCORD:
        return (
          <a
            className="text-app-foreground-inactive"
            href={helpDesks[type] ? helpDesks[type].link : ''}
            target={'_blank'}
          >
            {helpDesks[type] ? helpDesks[type].name : ''}
          </a>
        );
        break;
      default:
        return <span />;
    }
  };

  return (
    <div className="absolute top-0 left-0 right-0 bottom-0 flex items-center justify-center text-app-foreground">
      <div className="flex flex-col opacity-50">
        {Object.values(helpDesks).map((help, index) => {
          return (
            <div className=" text-app-foreground-inactive flex items-center mb-2 text-xl" key={`help-${client}-${index}`}>
              <div className="pr-1 flex items-center justify-center">
                {help.image}
              </div>
              {_renderLink(help.id)}
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default Help;