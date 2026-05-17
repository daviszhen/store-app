export interface Store {
  id: number;
  name: string;
  logo: string;
  theme: string;
  banner: string;
  notice: string;
  business_type: string;
}

export interface Category {
  id: number;
  name: string;
  icon: string;
  sort: number;
}

export interface Product {
  id: number;
  name: string;
  price: number;
  image: string;
  category_id: number;
  description: string;
  stock: number;
  status: number;
  category?: Category;
}

export interface CartItem {
  id: number;
  user_id: string;
  product_id: number;
  quantity: number;
  product?: Product;
}

export interface OrderItem {
  id: number;
  order_id: number;
  product_id: number;
  product_name: string;
  price: number;
  quantity: number;
}

export interface Order {
  id: number;
  order_no: string;
  user_id: string;
  total_amount: number;
  status: number;
  contact_name: string;
  contact_phone: string;
  contact_addr: string;
  remark: string;
  created_at: string;
  items?: OrderItem[];
}
