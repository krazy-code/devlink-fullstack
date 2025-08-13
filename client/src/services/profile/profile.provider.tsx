import type { Dispatch } from 'react';
import { createContext, useState } from 'react';
import type { ProfileItem } from './profile.types';

const ProfileContext = createContext<{
  profile: null | ProfileItem;
  setProfile: Dispatch<null | ProfileItem>;
}>({
  profile: null,
  setProfile: () => null,
});

interface ProfileProviderProps {
  children: React.ReactNode;
}

const ProfileProvider = ({ children }: ProfileProviderProps) => {
  const [profile, setProfile] = useState<ProfileItem | null>(null);

  return (
    <>
      <ProfileContext.Provider
        value={{
          profile,
          setProfile,
        }}
      >
        {children}
      </ProfileContext.Provider>
    </>
  );
};

export { ProfileContext, ProfileProvider };
