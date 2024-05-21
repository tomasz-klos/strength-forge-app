import { MenuItem } from "@atoms/menuItem";

import { HomeIcon } from "@assets/icons/home-icon";
import { SettingsIcon } from "@assets/icons/settings-icon";
import { UserIcon } from "@assets/icons/user-icon";

const Header: React.FC = () => {
  const menuItems = [
    { name: "Home", url: "/", icon: HomeIcon },
    { name: "Profile", url: "/profile", icon: UserIcon },
    { name: "Settings", url: "/settings", icon: SettingsIcon },
  ];

  return (
    <header className="sticky bottom-0 w-full bg-zinc-950 border-t border-zinc-800">
      <nav>
        <ul className="flex justify-end items-center gap-4 px-5">
          {menuItems.map((item) => (
            <MenuItem
              key={item.name}
              name={item.name}
              url={item.url}
              icon={item.icon}
            />
          ))}
        </ul>
      </nav>
    </header>
  );
};

export default Header;
