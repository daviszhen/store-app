import type { Store, Category, Product, CartItem, Order } from '../types';

const BASE_URL = 'http://localhost:8090/api';

async function request<T>(url: string, options?: RequestInit): Promise<T> {
  const res = await fetch(BASE_URL + url, {
    headers: { 'Content-Type': 'application/json' },
    ...options,
  });
  const json = await res.json();
  if (json.code !== 200) {
    throw new Error(json.msg || '请求失败');
  }
  return json.data as T;
}

export const api = {
  getStore: () => request<Store>('/store'),
  getCategories: () => request<Category[]>('/categories'),
  getProducts: (categoryId?: number) =>
    request<Product[]>(`/products${categoryId ? `?category_id=${categoryId}` : ''}`),
  getProduct: (id: number) => request<Product>(`/products/${id}`),

  getCart: (userId = 'default') => request<CartItem[]>(`/cart?user_id=${userId}`),
  addToCart: (productId: number, quantity = 1, userId = 'default') =>
    request<CartItem>('/cart', {
      method: 'POST',
      body: JSON.stringify({ user_id: userId, product_id: productId, quantity }),
    }),
  updateCartItem: (id: number, quantity: number) =>
    request<CartItem | null>(`/cart/${id}`, {
      method: 'PUT',
      body: JSON.stringify({ quantity }),
    }),
  deleteCartItem: (id: number) =>
    request<null>(`/cart/${id}`, { method: 'DELETE' }),

  createOrder: (contact: { name: string; phone: string; addr: string }, remark = '', userId = 'default') =>
    request<Order>('/orders', {
      method: 'POST',
      body: JSON.stringify({
        user_id: userId,
        contact_name: contact.name,
        contact_phone: contact.phone,
        contact_addr: contact.addr,
        remark,
      }),
    }),
  getOrders: (userId = 'default') => request<Order[]>(`/orders?user_id=${userId}`),
  getOrder: (id: number) => request<Order>(`/orders/${id}`),
};
