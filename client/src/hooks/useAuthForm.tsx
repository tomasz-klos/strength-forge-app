import { useNavigate } from "react-router-dom";
import { useForm, DefaultValues } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { ZodSchema } from "zod";
import { useMutation } from "@tanstack/react-query";

import { useToast } from "@atoms/use-toast";

interface UseAuthFormOptions<T> {
  schema: ZodSchema<T>;
  defaultValues: { [key in keyof T]: any };
  mutationFn: (data: T) => Promise<any>;
  onSuccessRedirect: string;
  successMessage: string;
  errorMessage: string;
}

const useAuthForm = <T extends Record<string, any>>({
  schema,
  defaultValues,
  mutationFn,
  onSuccessRedirect,
  successMessage,
  errorMessage,
}: UseAuthFormOptions<T>) => {
  const navigate = useNavigate();
  const { toast } = useToast();

  const form = useForm<T>({
    resolver: zodResolver(schema),
    defaultValues: defaultValues as DefaultValues<T>,
  });

  const mutation = useMutation({
    mutationFn,
    onSuccess: () => {
      navigate(onSuccessRedirect);
      toast({
        title: "Success",
        description: successMessage,
      });
    },
    onError: (error) => {
      console.error(error);
      toast({
        title: "Error",
        description: errorMessage,
        variant: "destructive",
      });
    },
  });

  const onSubmit = form.handleSubmit((data) => {
    mutation.mutate(data);
  });

  return { form, onSubmit };
};

export default useAuthForm;
