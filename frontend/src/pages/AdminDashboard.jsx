import { useState, useEffect } from 'react';
import api from '../api/axios';
import SweetForm from '../components/SweetForm';
import Transactions from '../components/Transactions';
import { Plus, Pencil, Trash2, Search, Package, History } from 'lucide-react';
import toast from 'react-hot-toast';

const AdminDashboard = () => {
    const [activeTab, setActiveTab] = useState('inventory');
    const [sweets, setSweets] = useState([]);
    const [loading, setLoading] = useState(true);
    const [showForm, setShowForm] = useState(false);
    const [editingSweet, setEditingSweet] = useState(null);
    const [searchQuery, setSearchQuery] = useState('');

    useEffect(() => {
        if (activeTab === 'inventory') {
            fetchSweets();
        }
    }, [activeTab]);

    const fetchSweets = async () => {
        try {
            const response = await api.get('/sweets');
            setSweets(response.data);
        } catch (error) {
            toast.error('Failed to fetch sweets');
        } finally {
            setLoading(false);
        }
    };

    const handleSave = async (sweetData) => {
        try {
            if (editingSweet) {
                await api.put(`/sweets/${editingSweet.ID}`, sweetData);
                toast.success('Sweet updated successfully');
            } else {
                await api.post('/sweets', sweetData);
                toast.success('Sweet added successfully');
            }
            setShowForm(false);
            setEditingSweet(null);
            fetchSweets();
        } catch (error) {
            toast.error(error.response?.data?.error || 'Operation failed');
        }
    };

    const handleDelete = async (id) => {
        if (window.confirm('Are you sure you want to delete this sweet?')) {
            try {
                await api.delete(`/sweets/${id}`);
                toast.success('Sweet deleted successfully');
                fetchSweets();
            } catch (error) {
                toast.error('Failed to delete sweet');
            }
        }
    };

    const handleRestock = async (id) => {
        const quantity = prompt('Enter quantity to add:');
        if (quantity && !isNaN(quantity)) {
            try {
                await api.post(`/sweets/${id}/restock`, { quantity: parseInt(quantity) });
                toast.success('Restock successful');
                fetchSweets();
            } catch (error) {
                toast.error('Restock failed');
            }
        }
    };

    const filteredSweets = sweets.filter(sweet =>
        sweet.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
        sweet.category.toLowerCase().includes(searchQuery.toLowerCase())
    );

    return (
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
            <div className="flex flex-col md:flex-row justify-between items-center mb-8 gap-4">
                <h1 className="text-3xl font-bold text-gray-900">Admin Dashboard</h1>
                {activeTab === 'inventory' && (
                    <button
                        onClick={() => {
                            setEditingSweet(null);
                            setShowForm(true);
                        }}
                        className="flex items-center bg-pink-600 text-white px-4 py-2 rounded-lg hover:bg-pink-700 transition-colors shadow-md"
                    >
                        <Plus className="h-5 w-5 mr-1" /> Add New Sweet
                    </button>
                )}
            </div>

            <div className="flex space-x-4 mb-6">
                <button
                    onClick={() => setActiveTab('inventory')}
                    className={`flex items-center px-4 py-2 rounded-lg font-medium transition-colors ${activeTab === 'inventory'
                            ? 'bg-pink-100 text-pink-700'
                            : 'bg-white text-gray-600 hover:bg-gray-50'
                        }`}
                >
                    <Package className="h-5 w-5 mr-2" />
                    Inventory
                </button>
                <button
                    onClick={() => setActiveTab('transactions')}
                    className={`flex items-center px-4 py-2 rounded-lg font-medium transition-colors ${activeTab === 'transactions'
                            ? 'bg-pink-100 text-pink-700'
                            : 'bg-white text-gray-600 hover:bg-gray-50'
                        }`}
                >
                    <History className="h-5 w-5 mr-2" />
                    Transactions
                </button>
            </div>

            {activeTab === 'inventory' ? (
                <div className="bg-white rounded-xl shadow-md overflow-hidden">
                    <div className="p-4 border-b border-gray-100 bg-gray-50 flex items-center">
                        <Search className="text-gray-400 h-5 w-5 mr-2" />
                        <input
                            type="text"
                            placeholder="Search inventory..."
                            value={searchQuery}
                            onChange={(e) => setSearchQuery(e.target.value)}
                            className="bg-transparent border-none focus:ring-0 w-full text-gray-700 outline-none"
                        />
                    </div>
                    <div className="overflow-x-auto">
                        <table className="min-w-full divide-y divide-gray-200">
                            <thead className="bg-gray-50">
                                <tr>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Category</th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Price</th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Stock</th>
                                    <th className="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
                                </tr>
                            </thead>
                            <tbody className="bg-white divide-y divide-gray-200">
                                {filteredSweets.map((sweet) => (
                                    <tr key={sweet.ID} className="hover:bg-gray-50 transition-colors">
                                        <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{sweet.name}</td>
                                        <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{sweet.category}</td>
                                        <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">${sweet.price.toFixed(2)}</td>
                                        <td className="px-6 py-4 whitespace-nowrap text-sm">
                                            <span className={`px-2 inline-flex text-xs leading-5 font-semibold rounded-full ${sweet.quantity > 10 ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                                                }`}>
                                                {sweet.quantity}
                                            </span>
                                            <button
                                                onClick={() => handleRestock(sweet.ID)}
                                                className="ml-2 text-xs text-blue-600 hover:text-blue-800 underline"
                                            >
                                                Restock
                                            </button>
                                        </td>
                                        <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                                            <button
                                                onClick={() => {
                                                    setEditingSweet(sweet);
                                                    setShowForm(true);
                                                }}
                                                className="text-indigo-600 hover:text-indigo-900 mr-4"
                                            >
                                                <Pencil className="h-5 w-5" />
                                            </button>
                                            <button
                                                onClick={() => handleDelete(sweet.ID)}
                                                className="text-red-600 hover:text-red-900"
                                            >
                                                <Trash2 className="h-5 w-5" />
                                            </button>
                                        </td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    </div>
                </div>
            ) : (
                <Transactions />
            )}

            {showForm && (
                <SweetForm
                    sweet={editingSweet}
                    onClose={() => {
                        setShowForm(false);
                        setEditingSweet(null);
                    }}
                    onSave={handleSave}
                />
            )}
        </div>
    );
};

export default AdminDashboard;
