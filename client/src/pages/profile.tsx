import { useOutletContext } from "react-router-dom";

import type { OutletContextValue } from "@shared/outletContext_types";

const Profile: React.FC = () => {
  const { user } = useOutletContext<OutletContextValue>();

  return (
    <section className="flex-1 flex flex-col gap-12">
      <h1 className="text-xl font-medium text-center">{user.name}</h1>
    </section>
  );
};

export default Profile;
