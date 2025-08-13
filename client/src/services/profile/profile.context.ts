import type { Dispatch } from 'react';
import { createContext } from 'react';
import type { ProfileItem } from './profile.types';

export const ProfileContext = createContext<{
  profile: null | ProfileItem;
  setProfile: Dispatch<null | ProfileItem>;
}>({
  profile: null,
  setProfile: () => null,
});
