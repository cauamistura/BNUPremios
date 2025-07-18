import React, { useState, useEffect } from 'react';
import type { Reward } from '../../Models/Reaward';
import './index.css';

interface RewardFormProps {
    reward?: Reward;
    onSubmit: (data: RewardFormData) => void;
    onCancel: () => void;
    loading?: boolean;
}

export interface RewardFormData {
    name: string;
    description: string;
    image: string;
    images: string[];
    draw_date: string; // Formato "YYYY-MM-DD" para o input, convertido para ISO no submit
    min_quota: number;
    price: number;
}

const RewardForm: React.FC<RewardFormProps> = ({ reward, onSubmit, onCancel, loading = false }) => {
    const [formData, setFormData] = useState<RewardFormData>({
        name: '',
        description: '',
        image: '',
        images: [''],
        draw_date: '',
        min_quota: 1,
        price: 0
    });
    const [errors, setErrors] = useState<Partial<RewardFormData>>({});

    useEffect(() => {
        if (reward) {
            setFormData({
                name: reward.name,
                description: reward.description,
                image: reward.image,
                images: reward.images || [''],
                draw_date: reward.draw_date.split('T')[0], // Converte ISO para formato date (YYYY-MM-DD)
                min_quota: reward.min_quota || 1,
                price: reward.price || 0
            });
        }
    }, [reward]);

    const validateForm = (): boolean => {
        const newErrors: Partial<RewardFormData> = {};

        if (!formData.name.trim()) {
            newErrors.name = 'Nome é obrigatório';
        }

        if (!formData.description.trim()) {
            newErrors.description = 'Descrição é obrigatória';
        }

        if (!formData.image.trim()) {
            newErrors.image = 'Imagem principal é obrigatória';
        }

        if (!formData.draw_date) {
            newErrors.draw_date = 'Data do sorteio é obrigatória';
        }

        setErrors(newErrors);
        return Object.keys(newErrors).length === 0;
    };

    const handleSubmit = () => {
        if (validateForm()) {
            // Converter a data para formato ISO
            const formDataWithISODate = {
                ...formData,
                draw_date: new Date(formData.draw_date + 'T00:00:00').toISOString()
            };
            onSubmit(formDataWithISODate);
        }
    };

    const handleInputChange = (field: keyof RewardFormData, value: string | number) => {
        setFormData(prev => ({ ...prev, [field]: value }));
        if (errors[field]) {
            setErrors(prev => ({ ...prev, [field]: undefined }));
        }
    };

    const addImage = () => {
        setFormData(prev => ({
            ...prev,
            images: [...prev.images, '']
        }));
    };

    const removeImage = (index: number) => {
        setFormData(prev => ({
            ...prev,
            images: prev.images.filter((_, i) => i !== index)
        }));
    };

    const updateImage = (index: number, value: string) => {
        setFormData(prev => ({
            ...prev,
            images: prev.images.map((img, i) => i === index ? value : img)
        }));
    };

    return (
        <div className="reward-form-page">
            <div className="reward-form-container">
                <h2 className="reward-form-title">
                    {reward ? 'Editar Prêmio' : 'Cadastrar Novo Prêmio'}
                </h2>
                
                <div className="reward-form">
                    <div className="form-group">
                        <label htmlFor="name">Nome do Prêmio *</label>
                        <input
                            type="text"
                            id="name"
                            value={formData.name}
                            onChange={(e) => handleInputChange('name', e.target.value)}
                            className={errors.name ? 'error' : ''}
                            placeholder="Digite o nome do prêmio"
                        />
                        {errors.name && <span className="error-message">{errors.name}</span>}
                    </div>

                    <div className="form-group">
                        <label htmlFor="description">Descrição *</label>
                        <textarea
                            id="description"
                            value={formData.description}
                            onChange={(e) => handleInputChange('description', e.target.value)}
                            className={errors.description ? 'error' : ''}
                            placeholder="Digite a descrição do prêmio"
                            rows={3}
                        />
                        {errors.description && <span className="error-message">{errors.description}</span>}
                    </div>

                    <div className="form-group">
                        <label htmlFor="image">Imagem Principal (URL) *</label>
                        <input
                            type="url"
                            id="image"
                            value={formData.image}
                            onChange={(e) => handleInputChange('image', e.target.value)}
                            className={errors.image ? 'error' : ''}
                            placeholder="https://exemplo.com/imagem.jpg"
                        />
                        {errors.image && <span className="error-message">{errors.image}</span>}
                    </div>

                    <div className="form-group">
                        <label>Imagens Adicionais (URLs)</label>
                        {formData.images.map((img, index) => (
                            <div key={index} className="image-input-group">
                                <input
                                    type="url"
                                    value={img}
                                    onChange={(e) => updateImage(index, e.target.value)}
                                    placeholder="https://exemplo.com/imagem.jpg"
                                />
                                {formData.images.length > 1 && (
                                    <button
                                        type="button"
                                        onClick={() => removeImage(index)}
                                        className="remove-image-btn"
                                    >
                                        Remover
                                    </button>
                                )}
                            </div>
                        ))}
                        <button
                            type="button"
                            onClick={addImage}
                            className="add-image-btn"
                        >
                            + Adicionar Imagem
                        </button>
                    </div>

                    <div className="form-row">
                        <div className="form-group">
                            <label htmlFor="draw_date">Data do Sorteio *</label>
                            <input
                                type="date"
                                id="draw_date"
                                value={formData.draw_date}
                                onChange={(e) => handleInputChange('draw_date', e.target.value)}
                                className={errors.draw_date ? 'error' : ''}
                            />
                            {errors.draw_date && <span className="error-message">{errors.draw_date}</span>}
                        </div>

                        <div className="form-group">
                            <label htmlFor="min_quota">Quantidade Mínima *</label>
                            <input
                                type="number"
                                id="min_quota"
                                value={formData.min_quota}
                                onChange={(e) => handleInputChange('min_quota', parseInt(e.target.value) || 1)}
                                className={errors.min_quota ? 'error' : ''}
                                min="1"
                            />
                            {errors.min_quota && <span className="error-message">{errors.min_quota}</span>}
                        </div>

                        <div className="form-group">
                            <label htmlFor="price">Preço (R$) *</label>
                            <input
                                type="number"
                                id="price"
                                value={formData.price}
                                onChange={(e) => handleInputChange('price', parseFloat(e.target.value) || 0)}
                                className={errors.price ? 'error' : ''}
                                min="0"
                                step="0.01"
                            />
                            {errors.price && <span className="error-message">{errors.price}</span>}
                        </div>
                    </div>

                    <div className="form-actions">
                        <button
                            type="button"
                            onClick={onCancel}
                            className="cancel-btn"
                            disabled={loading}
                        >
                            Cancelar
                        </button>
                        <button
                            type="button"
                            onClick={handleSubmit}
                            className="submit-btn"
                            disabled={loading}
                        >
                            {loading ? 'Salvando...' : (reward ? 'Atualizar' : 'Cadastrar')}
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default RewardForm; 