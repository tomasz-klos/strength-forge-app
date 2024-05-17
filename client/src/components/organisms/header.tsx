import { HomeIcon } from "@assets/icons/home-icon";
import { SettingsIcon } from "@assets/icons/settings-icon";
import { MenuItem } from "@atoms/menuItem";

const Header: React.FC = () => {
  const menuItems = [
    { name: "Home", url: "/", icon: HomeIcon },
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
