import { zodResolver } from "@hookform/resolvers/zod"
import { z } from "zod"
import { useForm } from "react-hook-form"
import { useCategories } from "../../../hooks/categories"
import { BaseFormProps } from "../../../types"
import { Button } from "../../../components/ui/button"
import { Input } from "../../../components/ui/input"
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "../../../components/ui/form"

interface UpdateCategoryFormProps extends BaseFormProps {
  id: string;
  name: string;
}

const formSchema = z.object({
  name: z.string().min(1, {
    message: "Name is required",
  }).max(150, {
    message: "Name is too long",
  }),
})

export function UpdateCategoryForm({ onClose, onRefresh, id, name }: UpdateCategoryFormProps) {
  const { updateCategory } = useCategories();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: name,
    },
  })

  async function onSubmit(values: z.infer<typeof formSchema>) {
    await updateCategory(id, values);
    onRefresh();
    if (onClose) { onClose() }
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Name</FormLabel>
              <FormControl>
                <Input placeholder={name}{...field} />
              </FormControl>
              <FormDescription>
                Category name.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />

        <div className="flex justify-between">
          <Button type="button" variant="ghost" onClick={onClose}>Close</Button>
          <Button type="submit">Update</Button>
        </div>
      </form>
    </Form>
  )
}