import { Autocomplete, AutocompleteProps, Chip, TextField, TextFieldProps } from "@mui/material";

interface Props extends Partial<AutocompleteProps<any, true, true, true>> {
  chipVariant?: "filled" | "outlined";
  textFieldProps?: Partial<TextFieldProps>;
}

const ChipInput = (props: Props) => {
  const { textFieldProps, chipVariant, ...autocompleteProps } = props;

  return (
    <Autocomplete
      {...autocompleteProps}
      multiple
      disableClearable
      options={[]}
      freeSolo
      renderTags={(value: any[], getTagProps: (arg0: { index: any }) => JSX.IntrinsicAttributes) =>
        value.map((option: any, index: any) => {
          return <Chip key={index} variant={chipVariant} size="small" label={option} {...getTagProps({ index })} />;
        })
      }
      renderInput={(params: any) => <TextField {...textFieldProps} {...params} />}
    />
  );
};

export default ChipInput;
