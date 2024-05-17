import { Link, useLocation } from "react-router-dom";
import { cn } from "@lib/utils";
import { IconProps } from "@assets/icons";

type MenuItemProps = {
  name: string;
  url: string;
  icon: React.FC<IconProps>;
};

export const MenuItem = ({ name, url, icon }: MenuItemProps) => {
  const location = useLocation();
  const isActive = location.pathname === url;

  const Icon = icon;

  return (
    <li className="flex-1">
      <Link
        to={url}
        className={cn(
          "flex flex-col items-center justify-center gap-1 px-3 py-2 hover:text-zinc-50 transition-colors duration-200 ease-in-out",
          isActive ? "text-zinc-50" : "text-zinc-500"
        )}
      >
        <Icon className="size-5" variant={isActive ? "filled" : "outline"} />
        <span className="text-xs">{name}</span>
      </Link>
    </li>
  );
};
