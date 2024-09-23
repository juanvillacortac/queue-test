export type Client = {
  id: number;
  dpi: string;
  name: string;
  clientType: number;
};

export type HistoryEntry = {
  id: number;
  client: Client;
  attendedBy: {
    id: number;
    email: string;
  };
  attendedAt: string;
  clientType: number;
};

export type QueueClient = Client & {
  arrival: string;
  requiredOperations: number;
};

export type Queue = {
  items: QueueClient[];
  capacity: number;
};

export function clientTypeString(number: number): string {
  return (
    {
      0: "VIP",
      1: "Regular",
    }[number] || "Desconocido"
  );
}
