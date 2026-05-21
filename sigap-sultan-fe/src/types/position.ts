export type Position = {
  id: number;
  clientId: number;
  name: string;
  privileges: Privilege[];
};

type Privilege = {
  menu: string;
  permissions: {
    read: number;
    create: number;
    update: number;
    delete: number;
  };
};
