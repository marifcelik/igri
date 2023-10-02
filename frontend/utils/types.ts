import type { Icon } from "lucide-vue-next";

export type DropdownItems = {
  icon: Icon;
  label: string;
  shortcut?: string;
  subMenu?: DropdownItems[];
  disabled?: boolean;
};