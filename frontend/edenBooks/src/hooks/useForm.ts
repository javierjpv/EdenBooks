import  { ChangeEvent, FormEvent, useState } from "react";

//Custom Hook para el manejo de formularios

export const useForm = <T extends Record<string, any>>(formulario: T) => {
  const [form, setForm] = useState<T>(formulario);

  const handleOnChange = (e: ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setForm({ ...form, [e.target.name]: value });
  };

  const handleSubmit = async (e: FormEvent<HTMLFormElement>, accionTrasSubmit: () => Promise<void>) => {
    e.preventDefault();
    await accionTrasSubmit();
  };

  const handleReset = () => {
    setForm(formulario);
  };

  return {
    ...form,
    form,
    handleOnChange,
    handleSubmit,
    handleReset,
  };
};
